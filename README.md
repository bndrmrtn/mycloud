# MyCloud

MyCloud is a cloud storage service that allows users to store and access files from anywhere. Users can upload files to the cloud, download files from the cloud, and delete files from the cloud. MyCloud is a web application that is accessible from any device with an internet connection.

## Installation

To install MyCloud, simply clone the repository and run the following command:

Clone the repository.
```bash
git clone https://github.com/bndrmrtn/mycloud.git
```

## Configuration

Under the backend directory, modify the `config.yaml` file to configure the application.

```yaml
service:
  version: "1"
  appdata_dir: "mycloud-appdata"
application:
  authorization:
    use_whitelist: false
    use_blacklist: false
    admin:
      primary_admin_email: "your@email.com" # Replace with your email
      enable_multi_admin: true
```

This configuration file allows you to configure the application version, the directory where the application data will be stored, and the authorization settings.
**NOTE: When both use_whitelist and use_blacklist are set to false, all users will be able to access the application.**

## Usage

Use docker compose to build and run the application.
```bash
docker compose up
```

To use MyCloud, simply navigate to `http://localhost:3000` in your web browser and you are ready to go. 🚀
