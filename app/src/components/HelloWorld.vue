<template>
  <div class="hello-world">
    <!-- Create an instance of the todo-item component -->
    <h1>{{ msg }}</h1>
    <h2>Hello World</h2>
    <h3>{{ msg2 }}</h3>
    <br>
    <md-avatar class="md-small"></md-avatar>
    <div>
      <md-checkbox v-model="checkboxValue.one" value="One">One</md-checkbox>
      <md-checkbox v-model="checkboxValue.two" value="TWO">Two</md-checkbox>

    </div>
    <div>
      <ol>
        <li> ONE: {{ checkboxValue.one }} </li>
        <li> TWO: {{ checkboxValue.two }} </li>
      </ol>
    </div>
    <md-button class="md-raised" v-on:click="greet">Handler Test </md-button>
    <md-input v-model="msg2" placeholder="edit me"></md-input>
  </div>
</template>

<script>
  import checkbox from 'vue-strap/src/Checkbox'

  export default {
    components: {checkbox},
    name: 'hello-world',
    data () {
      return {
        msg: 'Vue says Hello World',
        msg2: '',
        message: 'You loaded this page on ' + new Date().toLocaleString(),
        todos: [
          { text: 'Learn JavaScript' },
          { text: 'Learn Vue' },
          { text: 'Build something awesome' }
        ],
        checkboxValue: {
          'one': null,
          'two': 'two',
          'three': null,
          'four': null
        }
      }
    },
    // define methods under the `methods` object
    methods: {
      greet: function (event) {
        this.checkboxValue.one = 'ONE'
        this.msg2 = this.checkboxValue.one
      }
    },
    // lifecycle
    created: function () {
      // `this` points to the vm instance
      console.log('HelloWorld is: ' + this.msg)
    }
  }
</script>

<style scoped>

</style>
