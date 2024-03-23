/** @type {import('.types').PageServerLoad} */
export const load = ({ params }) => {
    const uuid = params['uuid'] as string
    const manifest_name = params['manifest_name'] as string

    return {
        uuid, manifest_name
    }
}

