import React from 'react'
import ReactDOM from 'react-dom'

import {BrowserRouter as Router,Route  } from 'react-router-dom'

import App from './App'


ReactDOM.render(
    <Router>
        <Route component={App}/>
    </Router>,
    document.querySelector('#root')
)