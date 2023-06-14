FROM python:3.10 AS build

# Copy the requirements file into the image
COPY ./requirements.txt /app/metrics/requirements.txt

# Switch working directory
WORKDIR /app/metrics

ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1

# Install the dependencies and packages in the requirements file
RUN pip install --no-cache-dir -r requirements.txt

# Copy every content from the local file to the image
COPY . /app/metrics

# Second stage: Runtime
FROM python:3.10-slim

# Switch working directory
WORKDIR /app/metrics

# Copy the requirements file from the build stage
COPY --from=build /app/metrics/requirements.txt /app/metrics/requirements.txt

# Install the packages specified in the requirements file
RUN pip install --no-cache-dir -r requirements.txt

# Copy only necessary files from the build stage
COPY --from=build /app/metrics /app/metrics

# Run the application
CMD ["python", "pol_client.py","&","python", "temp_client.py"]