import { gql } from 'urql';
import { pipe, subscribe } from 'wonka';
import { client } from './graphqlClient';

export interface Worktree {
  beanId: string;
  name: string | null;
  branch: string;
  path: string;
}

const WORKTREE_FIELDS = `
  beanId
  name
  branch
  path
`;

const WORKTREES_SUBSCRIPTION = gql`
  subscription WorktreesChanged {
    worktreesChanged {
      ${WORKTREE_FIELDS}
    }
  }
`;

const START_WORK = gql`
  mutation StartWork($beanId: ID!) {
    startWork(beanId: $beanId) {
      ${WORKTREE_FIELDS}
    }
  }
`;

const STOP_WORK = gql`
  mutation StopWork($beanId: ID!) {
    stopWork(beanId: $beanId)
  }
`;

const CREATE_WORKTREE = gql`
  mutation CreateWorktree($name: String!) {
    createWorktree(name: $name) {
      ${WORKTREE_FIELDS}
    }
  }
`;

const REMOVE_WORKTREE = gql`
  mutation RemoveWorktree($id: ID!) {
    removeWorktree(id: $id)
  }
`;

class WorktreeStore {
  worktrees = $state<Worktree[]>([]);
  loading = $state(false);
  error = $state<string | null>(null);

  #unsubscribe: (() => void) | null = null;

  subscribe(): void {
    if (this.#unsubscribe) return;

    const { unsubscribe } = pipe(
      client.subscription(WORKTREES_SUBSCRIPTION, {}),
      subscribe((result: { data?: { worktreesChanged?: Worktree[] }; error?: Error }) => {
        if (result.error) {
          console.error('Worktree subscription error:', result.error);
          this.error = result.error.message;
          return;
        }

        const wts = result.data?.worktreesChanged;
        if (wts) {
          this.worktrees = wts;
        }
      })
    );

    this.#unsubscribe = unsubscribe;
  }

  unsubscribe(): void {
    if (this.#unsubscribe) {
      this.#unsubscribe();
      this.#unsubscribe = null;
    }
  }

  async startWork(beanId: string): Promise<boolean> {
    this.loading = true;
    this.error = null;

    const result = await client.mutation(START_WORK, { beanId }).toPromise();

    this.loading = false;

    if (result.error) {
      this.error = result.error.message;
      return false;
    }

    return true;
  }

  async stopWork(beanId: string): Promise<boolean> {
    this.loading = true;
    this.error = null;

    const result = await client.mutation(STOP_WORK, { beanId }).toPromise();

    this.loading = false;

    if (result.error) {
      this.error = result.error.message;
      return false;
    }

    return true;
  }

  async createWorktree(name: string): Promise<Worktree | null> {
    this.loading = true;
    this.error = null;

    const result = await client.mutation(CREATE_WORKTREE, { name }).toPromise();

    this.loading = false;

    if (result.error) {
      this.error = result.error.message;
      return null;
    }

    return result.data?.createWorktree ?? null;
  }

  async removeWorktree(id: string): Promise<boolean> {
    this.loading = true;
    this.error = null;

    const result = await client.mutation(REMOVE_WORKTREE, { id }).toPromise();

    this.loading = false;

    if (result.error) {
      this.error = result.error.message;
      return false;
    }

    return true;
  }

  hasWorktree(beanId: string): boolean {
    return this.worktrees.some((wt) => wt.beanId === beanId);
  }
}

export const worktreeStore = new WorktreeStore();
