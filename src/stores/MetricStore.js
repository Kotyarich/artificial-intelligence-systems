import MetricService from "../services/MetricService";

class MetricStore {
  constructor(metric) {
    this.metric = metric;
    this.metricService = new MetricService();
  }

  status = 'initial';
  guns = [];
  metrics = [];

  getMetric = async () => {
    try {
      const data = await this.metricService.getMetric(this.metric);
      this.guns = data.guns;
      this.metrics = data.metrics;
      this.status = "ok"
    } catch (error) {
      this.status = 'error';
    }
  };
}

export default MetricStore;