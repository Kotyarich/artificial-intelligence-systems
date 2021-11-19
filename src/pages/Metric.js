import React from "react";
import "./Metric.css"

class Metric extends React.Component {

  constructor(props) {
    super(props);
    this.state = {guns: [], metricArray: []};
  }

  componentDidMount() {
    this.props.metricStore.getMetric().then(
      () => {
        this.setState({
          guns: this.props.metricStore.guns,
          metricArray: this.props.metricStore.metrics,
        })
      }
    );
  };

  render() {
    return (
      <div>
        {this.state.guns.map((gun, i) => {
          return <div className="metric-gun">
            <span className="metric-gun-index">{i}</span>
            <span className="metric-gun-model">{gun.model}</span>
          </div>
        })}
        <div className="metric-table">
          <div className="metric-row">
            <span className="metric-table-item">    </span>
            {this.state.guns.map((_, i) => {
              return <span className="metric-table-item">{i}</span>
            })}
          </div>
          {this.state.metricArray.map((line, i) => {
            return <div className="metric-row">
              <span className="metric-table-item">{i}</span>
              {line.map((metric, _) => {
                return <span className="metric-table-item">{metric}</span>
              })}
            </div>
          })}
        </div>
      </div>
    );
  }
}

export default Metric;