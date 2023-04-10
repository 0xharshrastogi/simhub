import { RealmType } from "../common";
import { Realm } from "../common/Realm";
import { ResourceBase } from "../common/ResourceBase";

const REALM_ID = 1;

export class MagnatesRealm extends Realm {
  realmType: RealmType = RealmType.Magnates;

  baseURL: string;

  constructor(baseURL: string) {
    super();
    this.baseURL = baseURL;
  }

  async fetchResourceList(): Promise<ResourceBase[]> {
    const uri = `${this.baseURL}/${REALM_ID}/resource/`;
    const response = await fetch(uri);
    return response.json();
  }
}
