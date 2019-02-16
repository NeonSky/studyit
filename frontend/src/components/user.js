var React = require('react')
var axios = require('axios')


module.exports = React.createClass({
    getInitialState: function() {
        return {
            firstName : "",
            lastName : "",
            avatarUrl : ""
        };
    },
    render: function() {
        let me = this;

        if(this.state.firstName == "") {
            axios.get('http://localhost:9000/users/' + this.props.match.params.mail)
                .then(function (response) {
                    let data = response.data;
                    if(data != null) {
                        me.setState({
                            firstName : data.firstName,
                            lastName : data.lastName,
                            avatarUrl : data.avatarUrl
                        });
                    }
                    console.log(data);
                })
                .catch(function (error) {
                    console.log(error);
                });
        }

        return (
            <div>
                <h1>This is a course page for {this.props.match.params.mail}</h1>
                <p>Name: {this.state.firstName} {this.state.lastName}</p>
                <img src={this.state.avatarUrl} />
            </div>
        )
    }
})
