const baseUrl = 'http://localhost:8000/api/v1';

class CompareService {
  getGuns = async () => {
    const url = baseUrl + '/guns';

    const headers = new Headers();
    const options = {
      method: 'GET',
      headers,
    };

    const request = new Request(url, options);
    const response = await fetch(request);
    return response.json();
  };

  getClosestToOne = async (gun) => {
    const url = baseUrl + `/closest/one?gun=${gun}`;

    const headers = new Headers();
    const options = {
      method: 'GET',
      headers,
    };

    const request = new Request(url, options);
    const response = await fetch(request);
    return response.json();
  };

  getClosestToN = async (guns) => {
    let url = baseUrl + '/closest/several?guns=' + guns[0];
    for (let i = 1; i < guns.length; i++) {
      url += '&guns=' + guns[i]
    }

    const headers = new Headers();
    const options = {
      method: 'GET',
      headers,
    };

    const request = new Request(url, options);
    const response = await fetch(request);
    return response.json();
  };

  getClosest = async (guns, dislikes) => {
    let url = baseUrl + '/closest?likes=' + guns[0];
    for (let i = 1; i < guns.length; i++) {
      url += '&likes=' + guns[i]
    }
    url += ''.concat(...dislikes.map((dislike) => '&dislike=' + dislike));

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

export default CompareService;