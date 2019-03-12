<template>
    <div class="container">
        <div class="row justify-content-center align-items-center vh-100">
            <div class="col-sm-11 col-md-6 card">
                <h1 class="h1 align-self-center">Welcome</h1>
                <form class="col-12" >
                    <div class="form-group">
                        <label >Username</label>
                        <input v-model="username" class="form-control align-self-center" type="text"/>
                    </div>

                    <div class="form-group">
                        <label>Password</label>
                        <input v-model="password" class="form-control align-self-center" type="password"/>
                    </div>

                    <button v-on:click.prevent="login" type="submit" class="btn btn-primary mb-3">Submit</button>
                </form>
            </div>
            <div>
        </div>
    </div>
    </div>
</template>

<script lang="ts">
  import {Vue, Component} from 'vue-property-decorator';

  @Component
  export default class LogIn extends Vue {
        username = '';
        password = '';

    // methods
        login () {
            fetch('http://localhost:8080/login', {
                method: 'POST',
                body: JSON.stringify({username: this.username, password: this.password}),
                headers: {
                    'Content-Type': 'text/plain'
                },
                credentials: 'include'
            }).then(res => res.json())
                .then(response => console.log('Success:', response))
                .catch(error => console.error('Error:', error));
        }
  };
</script>

<style scoped>

</style>
