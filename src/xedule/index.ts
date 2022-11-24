import Logger from '@root/logger';
import axios, { AxiosInstance } from 'axios';
import { Cache, CacheContainer } from 'node-ts-cache';
import { MemoryStorage } from 'node-ts-cache-storage-memory';
import XeduleAuthMonitor from './auth';

const xeduleCache = new CacheContainer(new MemoryStorage());

export default class Xedule {
  private http: AxiosInstance;
  private endpoint: string;
  private authCookie: string;
  private authMonitor: XeduleAuthMonitor;
  private static instance: Xedule;

  constructor(endpoint: string = process.env.XEDULE_ENDPOINT) {
    this.endpoint = endpoint;
    this.authMonitor = new XeduleAuthMonitor(endpoint);

    // Make axios use the auth cookie in every request
    this.http = axios.create({
      baseURL: `${endpoint}/api`,
      transformRequest: (_, headers) => {
        headers['Cookie'] = this.authCookie || '';
      },
    });
  }

  startAuthMonitor() {
    this.authMonitor.cookieStream.on('cookies', (cookies) => {
      // Get only the cookies for the endpoint
      const filteredCookies = cookies.filter((c) =>
        this.endpoint.includes(c.domain)
      );

      // Convert the cookies to a cookie string
      const cookieStringParts = filteredCookies.map(
        (c) => `${c.name}=${c.value}`
      );
      const cookieString = cookieStringParts.join('; ');
      this.authCookie = cookieString;
    });

    this.authMonitor.monitor();
  }

  @Cache(xeduleCache)
  async groups() {
    return this.http
      .get('/group')
      .then((res) => res.data.map((e) => e as XeduleModels.Group));
  }

  @Cache(xeduleCache)
  async organisationalUnit() {
    return this.http
      .get('/organisationalUnit')
      .then((res) => res.data.map((e) => e as XeduleModels.OrganisationalUnit));
  }

  // Singleton
  static getInstance() {
    if (!Xedule.instance) {
      Xedule.instance = new Xedule();
    }

    return Xedule.instance;
  }
}
