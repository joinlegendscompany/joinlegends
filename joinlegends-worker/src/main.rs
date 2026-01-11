mod mail;
mod processor;

use crate::mail::send_mail;
use crate::processor::{QueueMessage, process_with_progress};

use futures::future::pending;

#[tokio::main]
async fn main() -> redis::RedisResult<()> {
    let client = redis::Client::open("redis://127.0.0.1/")?;

    let video_con = client.get_multiplexed_async_connection().await?;
    let email_con = client.get_multiplexed_async_connection().await?;

    tokio::spawn(video_worker(video_con));
    tokio::spawn(mail_worker(email_con));

    futures::future::pending::<()>().await;
    Ok(())
}

async fn video_worker(mut con: redis::aio::MultiplexedConnection) -> redis::RedisResult<()> {
    loop {
        let (_queue, msg): (String, String) = redis::cmd("BRPOP")
            .arg("video_queue")
            .arg("0")
            .query_async(&mut con)
            .await?;

        let parsed: QueueMessage = serde_json::from_str(&msg).unwrap();

        // proces
        println!("processando video {}", parsed.id);
    }
}

async fn mail_worker(mut con: redis::aio::MultiplexedConnection) -> redis::RedisResult<()> {
    loop {
        let (_queue, msg): (String, String) = redis::cmd("BRPOP")
            .arg("video_queue")
            .arg(0)
            .query_async(&mut con)
            .await?;
        let parsed: QueueMessage = serde_json::from_str(&msg).unwrap();

        send_mail(
            parsed.payload["to"].as_str().unwrap(),
            parsed.payload["subject"].as_str().unwrap(),
            parsed.payload["body"].as_str().unwrap(),
        );
    }
}
