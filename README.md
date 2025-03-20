# 🚀 Logging Service with Automated S3 Upload

This project is a simple logging service built in Go, containerized using Docker, and designed to automatically upload logs to an AWS S3 bucket every time the container exits.

## ✅ Prerequisites

Ensure you have the following installed:
- **Docker**
- **AWS CLI** (Pre-installed in the Docker image)
- **Go** (If building locally)

---

## 🛠️ Step 1: Clone the Repository

```bash
git clone https://github.com/your-repo/go-log-service.git
cd go-log-service
```

---

## 📦 Step 2: Set Up AWS Credentials

Create a `.env` file in the project directory:

```bash
touch .env
```

Add your AWS credentials and region:

```env
AWS_ACCESS_KEY_ID=your-access-key
AWS_SECRET_ACCESS_KEY=your-secret-key
AWS_DEFAULT_REGION=your-region
```

> **Note:** Never commit `.env` files to version control. Keep them secure.

---

## 🏗️ Step 3: Build the Docker Image

Run the following command to build the Docker image:

```bash
sudo docker build -t logging-service .
```

---

## 🚀 Step 4: Run the Container

Start the container using the following command:

```bash
sudo docker run -d \
  -v $(pwd)/logs:/app/logs \
  --env-file .env \
  --name logging-service \
  logging-service
```

- **`-v $(pwd)/logs:/app/logs`**: Mounts a local folder to store logs.
- **`--env-file .env`**: Passes AWS credentials securely.
- **`--name logging-service`**: Names the container.

---

## 📄 Step 5: Monitor Logs

View container logs:

```bash
sudo docker logs logging-service
```

To see if the container is running:

```bash
sudo docker ps -a
```

---

## 📥 Step 6: Verify Logs on S3

Once the container exits, logs will automatically upload to your specified S3 bucket.

You can check if the logs are uploaded using:

```bash
aws s3 ls s3://your-s3-bucket-name/
```

---

## 🧹 Step 7: Clean Up

To remove the container:

```bash
sudo docker rm logging-service
```

To clean up unused Docker containers and images:

```bash
sudo docker container prune
sudo docker image prune
```

---

## 🚦 Additional Commands

- Restart the container:
  ```bash
  sudo docker start logging-service
  ```

- Stop the container:
  ```bash
  sudo docker stop logging-service
  ```

- Access the running container's shell:
  ```bash
  sudo docker exec -it logging-service /bin/bash
  ```

---

## 🚧 Troubleshooting

- **AWS Credentials Error:** Ensure the `.env` file contains valid AWS credentials and the correct region.
- **Logs Not Uploading:** Confirm network connectivity and correct S3 permissions.
- **Container Not Stopping Automatically:** Ensure the Dockerfile has the correct `timeout` value in the `CMD`.

---

## 🎉 Happy Logging!
