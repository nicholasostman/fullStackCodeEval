import '../styles/searchResults.css';
import React, { Component } from 'react';
import {List } from 'semantic-ui-react';
import PropTypes from 'prop-types';

export default class SearchResults extends Component {

	render() {
		let jokes = [];
		this.props.jokes.forEach((joke, index) => {
			jokes.push(<List.Item key={index}>{joke.Value}</List.Item>);
		});

		return (
			<List animated divided>
				{jokes}
			</List>
		);
	}
}

SearchResults.propTypes = {
	jokes: PropTypes.array.isRequired
};
