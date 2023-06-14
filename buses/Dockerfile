FROM python:3.10 AS build

# Copy the requirements file into the image
COPY ./requirements.txt /app/buses/requirements.txt

# Switch working directory
WORKDIR /app/buses

ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1

# Install the dependencies and packages in the requirements file
RUN pip install --no-cache-dir -r requirements.txt

# Copy every content from the local file to the image
COPY . /app/buses

# Second stage: Runtime
FROM python:3.10-slim

# Switch working directory
WORKDIR /app/buses

# Copy the requirements file from the build stage
COPY --from=build /app/buses/requirements.txt /app/buses/requirements.txt

# Install the packages specified in the requirements file
RUN pip install --no-cache-dir -r requirements.txt

# Copy only necessary files from the build stage
COPY --from=build /app/buses /app/buses

# Run the application
CMD ["python", "client.py"]