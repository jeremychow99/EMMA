IS213 ESD AY22/23 G4T7

SETUP STEPS:
====Prerequisites====
This are the versions of Docker and Node we used:
 - Docker - 20.10.22
 - Node - 19.6.1

====Configuring Backend====
Ensure that there is a clean environment, with no other docker containers. 
Also ensure that Kong and its relevant containers are new and dont contain existing configurations.

1. From root of project folder, open terminal and enter `docker compose up`.

2. Access http://localhost:1337 in the browser to create an admin user for Konga GUI.

Username: admin
Email:    <your email address>
Password: adminadmin

3. Sign in, and connect Konga to Kong by creating a new connection.
Name: default
Kong Admin URL: http://kong:8001

4. Click on Snapshots located on the sidebar
5. Select IMPORT FROM FILE and import ./utils/emma_snapshot.json
6. Click on DETAILS for the new snapshot created
7. Select RESTORE, tick all boxes, then click IMPORT OBJECTS

====NOTE====
If you encounter errors due to dangling containers, or docker using previously created containers (e.g. for Kong, Konga),
run the following to rebuild all images:
`docker compose down`
`docker compose build`
`docker compose up --remove-orphans -d`
 

====Starting the application in Browser====
1. Ensure backend docker containers are running and API Gateway is properly setup

2. From folder root, navigate to frontend and install dependencies by running the following:
- `cd frontend`
- `npm install`

3. Launch application 
run `npm run dev`

4. Access application in browser (default is port 5173)



REGARDING USER CREDENTIALS:

The following Admin account is available by default:
name: admin1
password: admin1

The following Technician accounts are available:
1. name: jeremy
password: jeremy

2. name: peopleschoice
password: peopleschoice

YOU MAY CREATE ADDITIONAL ACCOUNTS VIA USERS/AUTH MICROSERVICE: http://localhost:8000/api/v1/register
Refer to Auth/User MS API Docs for details and required JSON Body fields.
You may also create TECHNICIAN accounts via the frontend signup page on the application.
(ADMIN ROLE accounts can only be created via the API and not on frontend)