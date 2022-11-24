import Xedule from '@root/xedule';

export default class OrganisationalUnitController {
  static async index() {
    return Xedule.getInstance().organisationalUnit();
  }
}
