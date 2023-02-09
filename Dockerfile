FROM mongo:latest

# Set the working directory to /app
WORKDIR /app

# Copy the JSON schema file into the working directory
COPY schema.json /app/schema.json

# Run the MongoDB instance with the schema file
CMD mongoimport --host 127.0.0.1 --db test --collection products --type json --file /app/schema.json --jsonArray
