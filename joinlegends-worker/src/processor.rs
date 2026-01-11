use redis::AsyncCommands;
use serde::{Deserialize, Serialize};
use std::process::Stdio;
// use tokio::io::{AsyncBufReadExt, BufReader};
use tokio::io::{AsyncBufReadExt, BufReader};
use tokio::process::Command;

#[derive(Debug, Serialize, Deserialize)]
pub struct QueueMessage {
    pub id: String,
    pub r#type: String,
    pub payload: serde_json::Value,
}

pub async fn get_video_duration(input: &str) -> f64 {
    let output = Command::new("ffprobe")
        .args(&[
            "-v",
            "error",
            "-show_entries",
            "format=duration",
            "-of",
            "default=noprint_wrappers=1:nokey=1",
            input,
        ])
        .output()
        .await
        .expect("ffprobe failed");

    String::from_utf8_lossy(&output.stdout)
        .trim()
        .parse::<f64>()
        .unwrap()
}

pub async fn process_with_progress(
    con: &mut redis::aio::MultiplexedConnection,
    video_id: &str,
    input: &str,
    output: &str,
    watermark: &str,
) -> redis::RedisResult<()> {
    let total_duration = get_video_duration(input).await;
    let mut last_percent = 0;

    let mut child = Command::new("ffmpeg")
        .args(&[
            "-i",
            input,
            "-i",
            watermark,
            "-filter_complex",
            "overlay=10:10",
            "-progress",
            "pipe:1",
            "-nostats",
            output,
        ])
        .stdout(Stdio::piped())
        .spawn()
        .map_err(|e| {
            redis::RedisError::from((
                redis::ErrorKind::IoError,
                "failed to spawn ffmpeg",
                e.to_string(),
            ))
        })?;

    let stdout = child.stdout.take().ok_or_else(|| {
        redis::RedisError::from((redis::ErrorKind::IoError, "ffmpeg stdout not available"))
    })?;

    let mut reader = BufReader::new(stdout).lines();

    while let Some(line) = reader.next_line().await.map_err(|e| {
        redis::RedisError::from((
            redis::ErrorKind::IoError,
            "failed to read ffmpeg stdout",
            e.to_string(),
        ))
    })? {
        if let Some(ms) = line.strip_prefix("out_time_ms=") {
            let current_secs = ms.parse::<f64>().unwrap_or(0.0) / 1_000_000.0;

            if total_duration > 0.0 {
                let percent = ((current_secs / total_duration) * 100.0) as i32;
                let rounded = (percent / 5) * 5;

                if rounded > last_percent && rounded <= 100 {
                    last_percent = rounded;

                    let payload = format!("{}:{}", video_id, rounded);
                    let _: () = con.publish("videos_updates", payload).await?;
                }
            }
        }
    }

    let status = child.wait().await.map_err(|e| {
        redis::RedisError::from((
            redis::ErrorKind::IoError,
            "failed to wait ffmpeg process",
            e.to_string(),
        ))
    })?;

    if !status.success() {
        let _: () = con
            .publish("videos_updates", format!("{}:error", video_id))
            .await?;

        return Err(redis::RedisError::from((
            redis::ErrorKind::IoError,
            "ffmpeg exited with error",
        )));
    }

    let _: () = con
        .publish("videos_updates", format!("{}:100", video_id))
        .await?;

    Ok(())
}
