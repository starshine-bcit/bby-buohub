# bby-buohub

bby-buohub is a simple micro-services based containerized video streaming platform.

## Team members

Sasha, Dennis, Jose

## Componenets

### Docker-Compose

In order to run the project via docker-compose, some manual setup is required after cloning this repo.

1. Generate a new set of keys with genkey tool.
2. Move the newly generated `authjwt` and `authjwt.pub` into `bby-buohub/auth/keys`
3. Follow the directions in the [up readme](./up/README.md) to create a `bby-buohub/up/.env.production` file
4. It may? be necessary to do `npm install` and `npm run dev` once in the up project to generate the typescript definitions
5. Create a `.env` in the `bby-buohub` folder with the following template. These variables are used in the `docker-compose.yml` file

```bash
DB_PASSWORD=
SERVER_ENV=
DB_ROOT_PASSWORD=
```

Once those steps are complete you should be able to build and run the images!

### Auth

[See the readme here](./auth/README.md)

### Cdn

[See the readme here](./cdn/README.md)

### Up

[See the readme here](./up/README.md)

### Genkey

[See the readme here](./genkey/README.md)

## License

Unless otherwise specified in particular services, this project is licensed under the GNU Affero General Public License (AGPLv3). You can view the full text of the license in the LICENSE file or [here](https://www.gnu.org/licenses/agpl-3.0.txt).