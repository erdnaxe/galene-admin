/**
 * Galène-admin JavaScript library
 */

class GroupAPI {
    /**
     * Interact with Galène-admin group API.
     * 
     * @param {function} displayCallback Function to update HTML list of groups
     */
    constructor(displayCallback) {
        this.displayCallback = displayCallback;
    }

    // Get the list of groups and call function to display it
    reloadGroups() {
        fetch('/api/group', {
            method: "GET",
        }).then(r => r.json()).then(this.displayCallback);
    }

    // Create a group
    createGroup() {
        fetch(`/api/group`, {
            method: "POST",
            body: JSON.stringify(newGroup)
        }).then(reloadGroups);
    }

    // Update a group
    updateGroup(name, newGroup) {
        fetch(`/api/group/${name}`, {
            method: "PUT",
            body: JSON.stringify(newGroup)
        }).then(reloadGroups);
    }

    // Delete a group by name
    deleteGroup(name) {
        fetch(`/api/group/${name}`, {
            method: "DELETE",
        }).then(reloadGroups);
    }
}
