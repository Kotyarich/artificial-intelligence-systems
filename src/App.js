import './App.css';
import React from "react";
import {
  withRouter,
  Switch,
  Route
} from "react-router-dom";

import MetricsPage from "./pages/MetricsPage";
import Metric from "./pages/Metric";
import ComparationPage from "./pages/ComparationPage"
import MetricStore from "./stores/MetricStore";
import CompareStore from "./stores/CompareStore";

function App() {
  return (
    <div className="App">
      <Switch>
        <Route exact path="/">
          <ComparationPage compareStore={new CompareStore()}/>
        </Route>
        <Route path="/metrics">
          <MetricsPage/>
        </Route>
        <Route path="/euclid">
          <Metric metricStore={new MetricStore('euclid')}/>
        </Route>
        <Route path="/manhattan">
          <Metric metricStore={new MetricStore('city')}/>
        </Route>
        <Route path="/pearson">
          <Metric metricStore={new MetricStore('pearson')}/>
        </Route>
        <Route path="/tree">
          <Metric metricStore={new MetricStore('tree')}/>
        </Route>
      </Switch>
    </div>
  );
}

export default withRouter(App);
