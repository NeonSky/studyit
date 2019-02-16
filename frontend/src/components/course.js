var React = require('react')
var axios = require('axios')


module.exports = React.createClass({
    render: function() {
        axios.get('http://localhost:9000/courses/' + this.props.match.params.code)
            .then(function (response) {
                let data = response.data;
                console.log(data);
            })
            .catch(function (error) {
                console.log(error);
            });

        return (
            <div>
                <h1>This is a course page for {this.props.match.params.code}</h1>
            </div>
        )
    }
})
