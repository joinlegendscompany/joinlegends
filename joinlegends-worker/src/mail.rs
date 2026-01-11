use lettre::{
  message::Mailbox,
  transport::smtp::authentication::Credentials,
  Message, SmtpTransport, Transport,
};

pub fn send_mail(
  to: &str,
  subject: &str,
  body: &str,
) {
  println!("Sending mail to {} with subject {} and body {}", to, subject, body);
}