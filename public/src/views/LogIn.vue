<template>
  <div class="container">
    <div class="row justify-content-center align-items-center vh-100">
      <div class="col-sm-11 col-md-6 card">
        <h1 class="h1 align-self-center">Welcome</h1>
        <form class="col-12">
          <div class="form-group">
            <label for="username">Username</label>
            <input class="form-control align-self-center" id="username" type="text" placeholder="Username"/>
          </div>

          <div class="form-group">
            <label for="password">Password</label>
            <input class="form-control align-self-center" id="password" type="password"/>
          </div>

          <button type="submit" class="btn btn-primary mb-3">Submit</button>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    name: 'LogIn',
    data () {
      return {
        user: 'Flynn'
      };
    }
  };
</script>

<style scoped>

</style>
