import $ from 'jquery';
import { urlConstants } from '../common/Constants';

class SearchService {

	searchJokes(term) {
		return $.post({
			url: urlConstants.serverURL + urlConstants.jokeEndpoint + term,
			type: 'POST',
			data: term,
			dataType: 'text',
			success: function(data){
				return data;
			},
			error: function(data){
				throw new Error(data);
			}
		});
	}
}

const searchService = new SearchService();
export default searchService;