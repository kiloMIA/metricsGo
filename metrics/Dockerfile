# Use the official Python base image
FROM python:3.11

# Set the working directory inside the container
WORKDIR /app

# Copy the requirements.txt file
COPY requirements.txt .

# Install the Python dependencies
RUN pip install --no-cache-dir -r requirements.txt

# Copy the project files
COPY . .

# Expose the gRPC server port
EXPOSE 50052

# Run the application
CMD ["python", "temp_client.py"]