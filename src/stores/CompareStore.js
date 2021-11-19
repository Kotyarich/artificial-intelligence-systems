import CompareService from "../services/CompareService";

class CompareStore {
  constructor() {
    this.compareService = new CompareService();
  }

  status = 'initial';
  allGuns = [];
  closest = [];

  getGuns = async () => {
    if (this.allGuns === null) {
      return;
    }

    try {
      this.allGuns = await this.compareService.getGuns();
      this.status = "ok"
    } catch (error) {
      this.status = 'error';
    }
  };

  getClosestToOne = async (gun) => {
    try {
      this.closest = await this.compareService.getClosestToOne(gun);
      console.log('aaa')
      console.log(this.closest);
      this.status = "ok"
    } catch (error) {
      console.log(error);
      this.status = 'error';
    }
  };

  getClosestToN = async (guns) => {
    try {
      this.closest = await this.compareService.getClosestToN(guns);
      this.status = "ok"
    } catch (error) {
      this.status = 'error';
    }
  };

  getClosest = async (guns, dislikes) => {
    try {
      this.closest = await this.compareService.getClosest(guns, dislikes);
      this.status = "ok"
    } catch (error) {
      this.status = 'error';
    }
  };
}

export default CompareStore;