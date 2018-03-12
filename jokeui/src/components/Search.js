import '../styles/search.css';
import React, { Component } from 'react';
import { Container } from 'semantic-ui-react';
import SearchBox from './SearchBox';
import SearchResults from './SearchResults';

export default class Search extends Component {

	constructor(props) {
		super(props);
		this.state = {
			jokes : []
		};

		this.updateJokes = this.updateJokes.bind(this);
	}

	updateJokes(jokes) {
		this.setState({jokes : jokes});
	}

	render() {
		return (
			<div>
				<header className="search-header">
					<h1 className="search-title">Chuck Norris Jokes HQ</h1>
					<SearchBox searchCallback={this.updateJokes}/>
				</header>
				<Container className="bodyContainer">
					<SearchResults jokes={this.state.jokes}/>
				</Container>
			</div>
		);
	}
}
