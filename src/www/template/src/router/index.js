import Vue from 'vue'
import Router from 'vue-router'
import Home from '../view/home'
import HomeIndex from '../view/home/index'
import HomeRobot from '../view/home/robot'
import HomeSkill from '../view/home/skill'
import HomeAbility from '../view/home/ability'
import HomeTools from '../view/home/tools'

Vue.use(Router)

export default new Router({
  routes: [
      {
          path: '/',
          component: Home,
          children: [
              {
                  path: '/',
                  name: 'HomeIndex',
                  component: HomeIndex
              },
              {
                  path: '/robot',
                  name: 'HomeRobot',
                  component: HomeRobot
              },
              {
                  path: '/skill',
                  name: 'HomeSkill',
                  component: HomeSkill
              },
              {
                  path: '/ability',
                  name: 'HomeAbility',
                  component: HomeAbility
              },
              {
                  path: '/tools',
                  name: 'HomeTools',
                  component: HomeTools
              },
          ]
      }
  ]
})
