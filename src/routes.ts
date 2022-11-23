import { FastifyInstance } from 'fastify';
import GroupController from './controllers/group-controller';

export function registerRoutes(app: FastifyInstance) {
  app.get('/group', GroupController.index);
}
