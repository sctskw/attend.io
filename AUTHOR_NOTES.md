Hi! Welcome! Thank you for checking out my API. This is a WIP and 
was more done as a rapid prototype to see how quickly a robust API
can be created.

### RoadMap

- Fix CORS in Live Documentation so queries Execute
- Add Unit Tests to coincide with the current E2E Tests
- Add In-Memory Database to mock Database calls with Fixtures
- Add Authentication via API Tokens or GCP API Gateway
- Add a UI!


### To Improve

- Testing Testing Testing: Currently, the E2E tests run on the Live Database. This is 
  not good in practice, but works for now because the tests clean themselves up. Adding 
  some more service and database mocking would allow us to rid ourselves of relying on GCP 
  during the testing sequence.
    
- Concurrency: Currently, the operations are fairly primitive so there was not need to
  over-optimize with routines and parallel programming. I suspect as traffic grows, this
  should be considered for scaling, but being mounted on AppEngine we achieve a reasonable
  level of scale out-of-the-box.
  
- Database Upgrades: Currently, we're using a simple Firestore DB with just a couple of
  collections to handle Talks and Attendees. Using NoSQL is useful during prototyping so
  you don't get bogged down with writing Schemas. Given the level of relational data this
  API has the potential for though, it would probably be useful to incorporate an RDBMS.
  Ideally, we can use a hybrid approach thats geared toward mobile, but also something to 
  has good cold storage.
  
### Technical Decisions

##### GoLang
For API projects, I prefer GoLang since it has a strong built-in HTTP library and 
being able to add in concurrency requires little effort. The addition of a strongly 
typed language to any API saves a lot of time and headache to some of the comparative
non-statically typed languages in the field today

##### Swagger
In order to rapidly produce this API, I chose to use [go-swagger](https://github.com/go-swagger/go-swagger)
to help build out the boiler plate code that goes a long with standing up an API. This 
helped in dramatic fashion to provide good structure to the codebase and also comes
with generatable documentation OOTB which is a huge win since writing docs by hand can
be very time consuming. The `swagger.yml` can be updated to add new endpoints and their
associated models with ease!

##### GCP Firestore DB
This provided and easily scalable NoSQL solution to avoid burnout from writing 
and rewriting schemas. This also is a more mobile-centric approach since rendering
UI pages won't get bogged down by database joining. This should scale well for at
lease the first 100k Users.


##### GCP AppEngine
I love seeing my work in the wild! AppEngine is easy to mount projects on and provides
a VERY low configuration CI/CD pipeline. Being able to upgrade live code with a simple
push to the `main` branch in the project is hugely rewarding.

