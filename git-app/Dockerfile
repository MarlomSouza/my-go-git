# Build Stage
FROM node:18-alpine AS build
WORKDIR /app
COPY package*.json ./
RUN npm cache clean --force && npm install --legacy-peer-deps
COPY . .
RUN npm run build

# Production Stage
FROM nginx:stable-alpine AS production
# Copy the build output to the Nginx HTML directory
COPY --from=build /app/build /usr/share/nginx/html
EXPOSE 80
# Start Nginx server
CMD ["nginx", "-g", "daemon off;"]