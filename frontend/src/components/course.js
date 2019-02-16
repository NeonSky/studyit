var React = require('react')

module.exports = React.createClass({
    render: function() {
        return (
            <div>
                <h1>This is a course page for {this.props.match.params.code}</h1>
            </div>
        )
    }
})
