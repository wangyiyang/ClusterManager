from pecan import rest


class ClusterController(rest.RestController):
	
    @expose(generic=True, template='index.html')
    def index(self):
        return dict()