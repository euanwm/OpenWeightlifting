## OpenWeightlifting Frontend

### Overview

This repository contains the frontend code for the OpenWeightlifting application. This is a [Next.js](https://nextjs.org/) project bootstrapped with [`create-next-app`](https://github.com/vercel/next.js/tree/canary/packages/create-next-app).

### Structure

The frontend code is organized into the following directories:
* api: Contains the TypeScript code for fetching data from the OpenWeightlifting backend API
* components: Contains the reusable React components that make up the user interface of the application.
* config: Contains configuration files for the application, such as the fonts and site colors.
* layouts: Contains the layout components for the application, such as the header and footer.
* pages: Contains the React components for the individual pages of the application, such as the leaderboard page and the lifter page.
* styles: Contains the CSS files for the application.

### Entry Point

The entry point for the React application is the `App.tsx` file. This file contains the root component for the application.

### Serving the Application

The `index.tsx` file is the index file for the application. It serves the React application and the associated CSS and TypeScript files.


### Getting Started

To get started, clone the repository and install the dependencies:

```
git clone https://github.com/openweightlifting/frontend
cd frontend
npm install
```

To start the development server, run:

```bash
npm run dev
# or
yarn dev
```

Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.

You can start editing the page by modifying `pages/index.js`. The page auto-updates as you edit the file.

[API routes](https://nextjs.org/docs/api-routes/introduction) can be accessed on [http://localhost:3000/api/hello](http://localhost:3000/api/hello). This endpoint can be edited in `pages/api/hello.js`.

The `pages/api` directory is mapped to `/api/*`. Files in this directory are treated as [API routes](https://nextjs.org/docs/api-routes/introduction) instead of React pages.


### Conclusion

This repository contains the frontend code for the OpenWeightlifting application. The frontend is built using Next.js and React. The repository also includes all of the necessary files to build, deploy, and document the application.



## Learn More

To learn more about Next.js, take a look at the following resources:

- [Next.js Documentation](https://nextjs.org/docs) - learn about Next.js features and API.
- [Learn Next.js](https://nextjs.org/learn) - an interactive Next.js tutorial.

You can check out [the Next.js GitHub repository](https://github.com/vercel/next.js/) - your feedback and contributions are welcome!

## Deploy on Vercel

The easiest way to deploy your Next.js app is to use the [Vercel Platform](https://vercel.com/new?utm_medium=default-template&filter=next.js&utm_source=create-next-app&utm_campaign=create-next-app-readme) from the creators of Next.js.

Check out our [Next.js deployment documentation](https://nextjs.org/docs/deployment) for more details.
