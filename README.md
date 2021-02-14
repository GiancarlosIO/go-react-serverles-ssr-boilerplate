# go-react-serverles-ssr-boilerplate
A fullstack boilerplate that integrates golang + webpack + react + serverless ssr

It's a SPA that uses:
- Golang: httprouter, postgres, gorm, custom request logger
- UI: React
- styling: styled-components and tailwindcss
- Webpack: webpack-dev-server, postcss, file-loader, graphql-tag, babel-loader,
- Tooling: Babel, Eslint, prettier, husky
- Serverless Framework for the Server Side Rendering. Each page served by golang makes a request to a nodejs lambda to get the HTML and CSS before rendering the SPA.


## Running
Right now your need terminal to up services (will improve in the future):
1. Backend:
```bash
  $ go run server.go
```
2. Frontend:
```bash
  $ cd frontend
  $ npm install
  $ npm start
```
3. Serverless SSR:
```bash
  $ cd serverless-ssr
  $ npm install
  $ npm start
```
4. Go to http://localhost:8080

I'll use this in my next side projects. (wish me luck!)

