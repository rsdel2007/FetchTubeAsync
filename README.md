FetchTubeAsync is a Golang project that fetches the latest YouTube videos based on a predefined search query, stores the video data in a database, and provides APIs for accessing the stored video information.

## Features
- Asynchronous fetching of YouTube videos at intervals.
- Storing video data in a database with proper indexing.
- Paginated API to retrieve stored videos in reverse chronological order.
- Basic search API to search videos using their title and description.
- Dockerized for easy deployment and scalability.

## Prerequisites
- Golang installed on your machine.
- Docker installed for containerization.
- YouTube API Key - Obtain your API key from the YouTube API Console and set it in the environment variable API_KEY.


## Getting Started
Clone the repository:

```
git clone [https://github.com/rsdel2007/FetchTubeAsync](https://github.com/rsdel2007/FetchTubeAsync).git
cd FetchTubeAsync
#Set your YouTube API key in the environment variable:
export API_KEY=your_youtube_api_key

#Build and run the Docker container:
docker build -t FetchTubeAsync .
docker run -p 8081:8081 FetchTubeAsync
Access the API at http://localhost:8081.
```
API Endpoints
GET /videos: Retrieve stored video data in a paginated response sorted by published datetime.

Example: http://localhost:8081/videos?page=1&pageSize=10

GET /search?q=query&page=1&pageSize=10: Search stored videos using their title and description.

Example: http://localhost:8081/search?q=golang


