import { InMemoryCache, Reference, makeVar } from '@apollo/client';

function readCookie(name) {
    name += '=';
    for (var ca = document.cookie.split(/;\s*/), i = ca.length - 1; i >= 0; i--)
        if (!ca[i].indexOf(name))
            return ca[i].replace(name, '');
}

// Initializes to true if localStorage includes a 'token' key,
// false otherwise
export const isLoggedInVar = makeVar<boolean>(readCookie("user")!==undefined);

// Initializes to an empty array
// export const cartItemsVar = makeVar<string[]>([]);

