<template>
  <div id="app" class="md-layout-item">
    <!-- Create an instance of the todo-item component -->
    <h1>{{ msg }}</h1>
    <h2>Hello World</h2>
    <h3>{{ msg2 }}</h3>
    <br>
    <md-avatar class="md-small"></md-avatar>
    <div>
      <md-checkbox ref="one" v-model="checkboxValue.one" value="One">One</md-checkbox>
      <md-checkbox id="two" v-model="checkboxValue.two" value="TWO">Two</md-checkbox>

    </div>
    <div>
      <ol>
        <li> ONE: {{ checkboxValue.one }} </li>
        <li> TWO: {{ checkboxValue.two }} </li>
      </ol>
    </div>
    <md-button class="md-raised" v-on:click="greet">Handler Test </md-button>
    <md-field>
      <md-input v-model="msg2" placeholder="edit me"></md-input>
      <md-input type='number' v-model.number="c1"></md-input>
    </md-field>
  </div>
</template>

<script>
  import checkbox from 'vue-strap/src/Checkbox'
  import input from 'vue-strap/src/Input'

  export default {
    components: {checkbox, input},
    name: 'hello-world',
    data () {
      return {
        msg: 'Vue says Hello World',
        msg2: '',
        c1: 9999,
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
        this.checkboxValue.one = 'greeting'
        console.log('handler: ' + this.$el.querySelector('#two').value)
        this.$el.querySelector('#two').selected === true
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
