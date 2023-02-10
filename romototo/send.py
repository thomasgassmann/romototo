import smtplib
import email.message

def send_message(config, subject, message):
    username = config['mail']['username']
    password = config['mail']['password']
    recipient = config['recipient']

    host = config['mail']['host']
    port = config['mail']['port']

    mail = email.message.Message()
    mail['Subject'] = subject
    mail['From'] = username
    mail['To'] = recipient
    mail.add_header('Content-Type', 'text/html')
    mail.set_payload(message, 'utf-8')

    server = smtplib.SMTP(f'{host}:{port}')
    server.ehlo()
    server.starttls()

    server.login(username, password)
    server.sendmail(username, recipient, mail.as_string())
    server.quit()
