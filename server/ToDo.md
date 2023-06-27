# MiddleWare
----------
## configure MiddleWares
- [ ] configure middleware in a structured map 
```Go
   _middlewares:=map[string]func(c *fiber.Ctx) error
```
- [ ] read routes from `config.yaml` file and apply relevant middleware to the route 
- [ ] middlewares to implement:
  - [ ] login __should return a `jwt` token upon success__
  - [ ] isLoggedIn
  - [ ] isEnabled
  - [ ] requiresLogin
  - [ ] getQueryParams *{should be enabled for all requests}*
  - [ ] getURLParams *{should be enabled for all requests}*
  - [ ] getPostParams *{should be enabled for all **post** requests}*