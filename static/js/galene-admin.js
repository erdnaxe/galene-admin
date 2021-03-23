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
        return fetch('/api/group', {
            method: "GET",
        }).then(r => r.json()).then(this.displayCallback);
    }

    // Create a group
    async createGroup(newGroup) {
        await fetch(`/api/group`, {
            method: "POST",
            body: JSON.stringify(newGroup)
        });
        this.reloadGroups();
    }

    // Update a group
    async updateGroup(name, newGroup) {
        await fetch(`/api/group/${name}`, {
            method: "PUT",
            body: JSON.stringify(newGroup)
        });
        this.reloadGroups();
    }

    // Delete a group by name
    async deleteGroup(name) {
        await fetch(`/api/group/${name}`, {
            method: "DELETE",
        });
        this.reloadGroups();
    }
}
