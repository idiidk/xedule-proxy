import Xedule from '@root/xedule';

export default class GroupController {
  static async index() {
    return Xedule.getInstance().groups();
  }
}
