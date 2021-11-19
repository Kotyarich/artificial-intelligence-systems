const baseUrl = 'http://localhost:8000/api/v1/metric';

class ForumService {
  getMetric = async (metricName) => {
    const url = baseUrl + '/' + metricName;

    const headers = new Headers();
    const options = {
      method: 'GET',
      headers,
    };

    const request = new Request(url, options);
    const response = await fetch(request);
    return response.json();
  };
}

export default ForumService;