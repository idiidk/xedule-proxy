import { FastifyInstance } from 'fastify';
import GroupController from './controllers/group-controller';
import OrganisationalUnitController from './controllers/organisational-unit-controller';

export function registerRoutes(app: FastifyInstance) {
  app.get('/group', GroupController.index);
  app.get('/organisationalUnit', OrganisationalUnitController.index);
}
