import React from "react";

class MetricsPage extends React.Component {

  render() {
    return(
      <div className={'main-page'}>
        <div className={'metric-list'}><a href={'euclid'}>Евклидово расстояние</a></div>
        <div className={'metric-list'}><a href={'manhattan'}>Манхеттенсое расстояние</a></div>
        <div className={'metric-list'}><a href={'pearson'}>Коррелияция Пирсона</a></div>
        <div className={'metric-list'}><a href={'tree'}>Древесная метрика</a></div>
      </div>
    )
  }
}

export default MetricsPage;