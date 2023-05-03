docker run -d \
  --name redis \
  -p 6379:6379 \
  -p 8001:8001 \
  -v "$(pwd)/data:/data" \
  redis/redis-stack:latest