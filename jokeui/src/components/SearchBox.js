import '../styles/App.css';
import '../styles/searchBox.css';
import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { Search } from 'semantic-ui-react';
import SearchService from '../Services/SeachService';

const placeHolderText = 'Search by keyword...';
const minChars = 3;
const timeoutDelay = 500;

export default class SearchBox extends Component {

	constructor(props) {
		super(props);

		this.state = {
			isLoading: false,
			results: [],
			value: ''
		};

		this.clearSearch = this.clearSearch.bind(this);
		this.handleSearchChange = this.handleSearchChange.bind(this);
	}

	clearSearch() {
		this.setState({
			isLoading: false,
			results: [],
			value: ''
		}, this.props.searchCallback([]));
	}

	handleSearchChange(e, { value }) {
		this.setState({ isLoading: true, value });

		if (this.lastRequestId !== null){
			clearTimeout(this.lastRequestId);
		}

		this.lastRequestId = setTimeout(() => {
			if (this.state.value.length < 1) return this.clearSearch();
			if (this.state.value.length > 2) { // API only takes 3 chars or more
				const requestPromise = SearchService.searchJokes(this.state.value);
				Promise.resolve(requestPromise).then((data) => {
					const jokes = JSON.parse(data).Result;
					this.setState({
						results: jokes,
						isLoading: false
					}, this.props.searchCallback(jokes));
				}).catch((error) => {
					this.setState({
						isLoading: false
					});
					throw error;
				});
			}
		}, timeoutDelay);

	}

	render() {
		const { isLoading, value } = this.state;
		this.placeHolderText = placeHolderText;

		const helpText = (this.state.value.length > 0 && this.state.value.length < minChars) ? '' : 'hidden';

		return (
			<div>
				<Search
					className="searchInputBox"
					alt="search box"
					placeholder={this.placeHolderText}
					fluid
					minCharacters={minChars}
					loading={isLoading}
					size="huge"
					open={false}
					onSearchChange={this.handleSearchChange}
					value={value}
				/>
				{<div className={helpText}>Enter at least 3 letters</div>}
			</div>
		);
	}
}

SearchBox.propTypes = {
	searchCallback: PropTypes.func.isRequired
};
