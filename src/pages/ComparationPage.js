import React from "react";
import "./Metric.css"

class ComparationPage extends React.Component {

  constructor(props) {
    super(props);
    this.state = {guns: [], closest: [], likes: '', dislikes: ''};
  }

  componentDidMount() {
    this.props.compareStore.getGuns().then(
      () => {
        this.setState({
          guns: this.props.compareStore.allGuns,
        })
      }
    );
  };

  closestToOne = (event) => {
    const i = +this.state.likes;
    this.props.compareStore.getClosestToOne(this.state.guns[i].ID).then(
      () => {
        this.setState({
          closest: this.props.compareStore.closest
        })
      }
    );
    event.preventDefault();
  };

  closestToN = (event) => {
    const numbers = this.state.likes.split(' ');
    const ids = numbers.map((number) => this.state.guns[number].ID);
    this.props.compareStore.getClosestToN(ids).then(
      () => {
        this.setState({
          closest: this.props.compareStore.closest
        })
      }
    );
    event.preventDefault();
  };

  closest = (event) => {
    const numbers = this.state.likes.split(' ');
    const idsLikes = numbers.map((number) => this.state.guns[number].ID);
    const dislikes = this.state.dislikes.split(' ');
    const idsDislikes = dislikes.map((number) => this.state.guns[number].ID);
    this.props.compareStore.getClosest(idsLikes, idsDislikes).then(
      () => {
        this.setState({
          closest: this.props.compareStore.closest
        })
      }
    );
    event.preventDefault();
  };

  likesOnChange = (event) => {
    this.setState({likes: event.target.value})
  };

  dislikesOnChange = (event) => {
    this.setState({dislikes: event.target.value})
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
        <div className="input">
          <div className="input-likes">
            Понравившиеся:
            <input value={this.state.likes} onChange={this.likesOnChange}/>
          </div>
          <div className="input-dislikes">
            Дизлайки:
            <input value={this.state.dislikes} onChange={this.dislikesOnChange}/>
          </div>
        </div>
        <div className="buttons">
          <button onClick={this.closestToOne}>Похожие на 1</button>
          <button onClick={this.closestToN}>Похожие на N</button>
          <button onClick={this.closest}>Похожие на N с дизлайками</button>
        </div>
        <div className="closest">
          <div className="closest-header">
            <b>Похожие:</b>
          </div>
          {this.state.closest.map((gun, i) => {
            return <div className="metric-gun">
              <span className="metric-gun-index">{i}</span>
              <span className="metric-gun-model">{gun.model}</span>
            </div>
          })}
        </div>
      </div>
    );
  }
}

export default ComparationPage;