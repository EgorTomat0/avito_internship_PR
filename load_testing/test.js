import http from 'k6/http';
import {check, sleep} from 'k6';
import {randomIntBetween, randomString} from 'https://jslib.k6.io/k6-utils/1.2.0/index.js';

export const options = {
    stages: [
        {duration: '1m', target: 100},
        {duration: '1m', target: 100},
        {duration: '2m', target: 300},
        {duration: '1m', target: 150},
    ],
    thresholds: {
        http_req_failed: ['rate<0.01'],
        http_req_duration: ['p(95)<301'],
        checks: ['rate>0.95'],
    },
};

const BASE_URL = 'http://localhost:8080';

function generateValidUUIDv4() {
    const bytes = new Uint8Array(16);
    crypto.getRandomValues(bytes);
    bytes[6] = (bytes[6] & 0x0f) | 0x40;
    bytes[8] = (bytes[8] & 0x3f) | 0x80;

    return Array.from(bytes)
        .map((b, i) => {
            const hex = b.toString(16).padStart(2, '0');
            if (i === 4 || i === 6 || i === 8 || i === 10) {
                return '-' + hex;
            }
            return hex;
        })
        .join('');
}

function generateTeamData() {
    const teamId = randomString(8).toLowerCase();
    const numMembers = randomIntBetween(1, 6);

    const members = [];
    for (let i = 0; i < numMembers; i++) {
        members.push({
            user_id: generateValidUUIDv4(),
            username: `user_${randomString(8).toLowerCase()}_${Date.now()}_${i}`,
            is_active: Math.random() > 0.1
        });
    }

    return {
        team_name: `team-${teamId}-${Date.now()}`,
        members: members
    };
}

const existingTeams = new Set();

export default function () {
    const teamData = generateTeamData();
    const payload = JSON.stringify(teamData);
    const params = {
        headers: {
            'Content-Type': 'application/json',
        },
        tags: {name: 'CreateTeam'},
    };
    const response = http.post(`${BASE_URL}/team/add`, payload, params);
    const checks = check(response, {
        'status is 201': (r) => r.status === 201,
        'response has team data': (r) => r.json('team') !== undefined,
        'team name matches': (r) => {
            const team = r.json('team');
            return team && team.team_name === teamData.team_name;
        },
        'response time acceptable': (r) => r.timings.duration < 2000,
        'has correct content type': (r) => r.headers['Content-Type']?.includes('application/json'),
    });
    if (response.status === 201) {
        check(response, {
            'team structure is correct': (r) => {
                const team = r.json('team');
                return team &&
                    team.team_name &&
                    team.members &&
                    Array.isArray(team.members) &&
                    team.members.length > 0;
            },
            'members have UUID v4 format': (r) => {
                const team = r.json('team');
                if (!team || !team.members) return false;
                return team.members.every(member => {
                    if (!member.user_id) return false;
                    const uuidRegex = /^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$/i;
                    return uuidRegex.test(member.user_id);
                });
            },
            'members have required fields': (r) => {
                const team = r.json('team');
                if (!team || !team.members) return false;

                return team.members.every(member =>
                    member.user_id &&
                    member.username &&
                    typeof member.is_active === 'boolean'
                );
            },
        });
        existingTeams.add(teamData.team_name);
    }
    if (response.status === 400) {
        check(response, {
            'error structure correct': (r) => {
                const error = r.json('error');
                return error && error.code && error.message;
            },
            'error code is TEAM_EXISTS': (r) => r.json('error.code') === 'TEAM_EXISTS',
        });
    }
    sleep(randomIntBetween(1, 3));
}