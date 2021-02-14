# go-react-serverles-ssr-boilerplate
A fullstack boilerplate that integrates golang + webpack + react + serverless ssr

It's a SPA that uses:
- Golang: httprouter, postgres, gorm
- UI: React
- styling: styled-components and tailwindcss
- Webpack: webpack-dev-server, postcss, file-loader, graphql-tag, babel-loader,
- Tooling: Babel, Eslint, prettier, husky
- Serverless Framework for the Server Side Rendering. Each page served by golang makes a request to a nodejs lambda to get the HTML and CSS before rendering the SPA.


## running
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
1. Backend:
```bash
  $ cd serverless-ssr
  $ npm install
  $ npm start
```

I'll use this in my next side projects. (wish me luck!)

