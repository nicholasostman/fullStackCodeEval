import './styles/App.css';
import React, { Component } from 'react';
import Search from './components/Search';

class App extends Component {
	render() {
		return (
			<div className="App">
				<Search />
			</div>
		);
	}
}

export default App;
