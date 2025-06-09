# Use official Python image
FROM python:3.11-slim

# Set working directory
WORKDIR /app

# Copy the app code
COPY . .

# Install dependencies
RUN pip install --no-cache-dir -r requirements.txt

# Expose the port the app runs on
EXPOSE 5000

# Run the application
CMD ["python", "app.py"]
