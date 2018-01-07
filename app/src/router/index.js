import Hello from '@/components/Hello'
import HelloWorld from '@/components/HelloWorld'

const routes = [
  {
    path: '/',
    name: 'Hello',
    component: Hello
  },
  {
    path: '/hw',
    name: 'Hello World',
    component: HelloWorld
  }
]
export default routes
