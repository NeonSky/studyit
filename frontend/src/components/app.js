var React = require('react')
var ReactDOM = require('react-dom')
const Route = require('react-router').Route
const IndexRoute = require('react-router').IndexRoute
const BrowserRouter = require('react-router-dom').BrowserRouter
const Switch = require('react-router-dom').Switch

var Hello = require('./hello')
var Course = require('./course')


module.exports = React.createClass({

    render: function() {
        return (
            <BrowserRouter>
                <Switch>
                    <Route path="/course/:code" component={Course} />
                    <Route path="/" component={Hello} />
                </Switch>
            </BrowserRouter>
        )
    }
})
