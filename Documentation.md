#Adalpha Test Documentation

###By Flynn Davies

#### 05/03/19
  - I Have created the basic folder structure
  - created the git repository (https://github.com/shik4rishuffle/adalpha
.git)
  - installed vue-cli with the webpack template including Jest testing, dev and prod build processes, and eslint.
  - began configuring the eslint to keep my codebase up to a good standard. found this pretty helpful guide on best 
  practices and style guides with Vue.js (https://itnext.io/how-to-structure-a-vue-js-project-29e4ddc1aeeb)
  - configured the webpack config to support webpack 4 and build for vue
  
  ### 06/03/19
  - Add in estimated views and components
    - I have only added in 4 pages. the app should not be overly complicated and easy to use so I want to show this by keeping it minimalistic
    - I have only added in a couple of the obvious components for now. due to my limited knowledge of using vue with GO I am going to add more as they become necessary.
  - Add in view to Router  
  - Add bootstrap to reduce time spent on styling and make the mobile first approach easier
  - Begin adding in basic DOM elements to Views
  
  ### 08/03/19
  - Add in TypeScript
  - Reconfigure existing Vue files to be Vue Class components with TypeScript
  - add in TypeScript basic Types Decelerations
  - wrap App in Declared Vue app to include all Vue functionality for TypeScript
  
  ### 10/03/19
  - Installed Docker to Computer
  - installed GO 
  - add in CORs to main.go to keep my projects clean and in separate front end/ back end folders (check with ad alpha to make sure if this is correct)
  - add in vue-resource for http requests
  - username is `ADA123456789` password is anything
  - got the login page returning the cookies
  
  worked out the logistics of ajax forms with go databases and vue with a little help from Matt. I had to play around
   with a few methods and plugins such as axios and vue-resource but finally settled on a simple fetch method. the 
   reason for this being that axios had a few issues with it for cors requests and vue-resources had since become 
   outdated and i would much rather go without plugins where possible to save space and dependencies.
  
  ### 12/03/19
  - add in front-authentication method for navigation
  - add in cookie functionality for staying logged in
  - handle log out functionality
  
  @TODO: fix the redirecting issue with cookies
  
  ### 14/03/19
  - handle login success/fails
  - pull through data
    - Going to attempt to just GET from database and map it to a table. this should work!
    - turned out to be a LOT fiddlier than I thought. It required breaking up the data and remapping the array of 
    objects so they all have accessible property names.
  - assuming that the unscaled amount and exponent is the amount of shares in each company and the total is 200,000 I
   will need to work out the scaled amount to create an easier to read graph. will need to check with ad-alpha that 
   this is correct
   - now feed through ISIN, current price and current percentage of portfolio to be kept in the vuex store.

