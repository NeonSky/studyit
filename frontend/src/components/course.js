var React = require('react')
var axios = require('axios')


module.exports = React.createClass({
    getInitialState: function() {
        return {
            code : "",
            description : "",
            level : "",
            name : ""
        };
    },
    render: function() {
        var me = this;
        if(this.state.code == "") {
            axios.get('http://localhost:9000/courses/' + this.props.match.params.code)
                .then(function (response) {
                    let data = response.data;
                    console.log(me);
                    console.log(data);
                    if(data != null) {
                        me.setState({
                            code : data.code,
                            description : data.description,
                            level : data.level,
                            name : data.name
                        });
                    }
                })
                .catch(function (error) {
                    console.log(error);
                });
        }

        return (
            <div>
                <h1>This is a course page for {this.props.match.params.code}</h1>
                <p>Course Code: {this.state.code}</p>
                <p>Description: {this.state.description}</p>
                <p>Level: {this.state.level}</p>
                <p>Name: {this.state.name}</p>
            </div>
        )
    }
})
