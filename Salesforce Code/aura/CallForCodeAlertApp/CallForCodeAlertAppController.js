({
    afterScriptsLoaded: function(component,event,helper){

    },
    deleteRec: function(component,event,helper){
        swal({
            title: 'Are you sure?',
            text: "You won't be able to revert this!",
            type: 'warning',
            showCancelButton: true,
            confirmButtonColor: '#3085d6',
            cancelButtonColor: '#d33',
            confirmButtonText: 'Yes, delete it!'
        }).then(function (){
            // This block of code gets executed if user chooses ok for deleting record.
            console.log("User chose to delete the record");
            // Here you can call the helper function.
        });
        ItemSet.prototype.itemFromElement = function(element) {
    var hasOwn = Object.prototype.hasOwnProperty.call;
    var cur = element;
    while (cur) {
        if (hasOwn(cur, 'delete-btn')) {
            return cur['delete-btn'];
        }
        cur = cur.parentNode;
    }

    return null;
};
        
    }
})